package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//Usecase
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	//Essa função vai ser um método POST da nossa API e vamos receber no body da requisição qual é o produto que
	//queremos inserir no banco de dados então precisamos de uma maneira para fazer com que o JSON enviado na rota
	//vire a estrutura product que temos aqui na nossa API
	var product model.Product

	//Essa função vai pegar o JSON que estamos recebendo na requisição e transformar ela no product que é a estrutura
	//que temos aqui dentro da nossa aplicação
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	//Se não der erro vou chamar o usecase que vai chamar o repository e o repository vai inserir os dados no banco
	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	//Nosso método espera receber um id_product e esse ID vai vir nos parâmetros da rota, para isso precisamos de uma
	//maneira de extraí-lo
	id := ctx.Param("productId")

	//Precisamos fazer validações para ver se o nosso parâmetro veio certo
	//Essa nossa função retorna uma string então precisamos fazer algumas verificações para ver se essa string não está
	//nula e se a gente consegue converter essa string em um inteiro
	if id == "" {
		response := model.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
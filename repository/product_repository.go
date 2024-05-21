package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	//Definindo a query de consulta no banco
	query := "SELECT id, product_name, price FROM product"

	//Conectando ao banco e executando a consulta
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	//Definindo a nossa lista para podermos jogar os itens retornados do banco
	var productList []model.Product
	var productObj model.Product

	//Percorrendo cada linha e adicionando o item a nossa lista
	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	//Retornando a lista
	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	//Criando a variável ID
	var id int

	//Escrevendo a query que quero executar no banco
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	//Executar a query de fato passando os parâmetros para ela
	//Como a query nos retorna um 'ID' podemos dar um Scan() para receber esse 'ID'
	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
	//Inicializando a query, preparando a query
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//Definindo a variável product que vai ser retornado pelo banco
	var product model.Product

	//Método para executar de fato a query e retornar os dados na nossa variável
	err = query.QueryRow(id_product).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		//Aqui queremos fazer algumas verificações porque quero verificar se o banco de dados não pode encontrar
		//nenhuma linha com esse ID que passei
		//Essa verificação significa que o banco não pode encontrar nenhum registro com o ID que passei
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &product, nil
}
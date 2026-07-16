# API de Productos

Este proyecto es una prueba tecnica que consiste en una API para gestionar productos, construida con Go, Graphql y Docker, siguiendo Clean Architecture, buenas practicas y los principios SOLID, la finalidad de esto es que nuestro codigo sea escalable, mantenible y testeable, dado que cuando nuestras capas estan fuertemente acopladas es dificil realizar test unitarios y realizar modificaciones, por ejemplo si quisieramos cambiar la base de datos, o quisieramos cambiar de Graphql a REST con esta estructura no tendriamos que realizar grandes modificaciones.

Este ejemplo contiene dos elementos precargados en el repositorio de productos para pruebas

### Nota: Los test se ejecutan automanticamente cuando se pone a correr el contenedor de docker

## Como ejecutar el proyecto
Para ejecutar este proyecto solo necesitas tener instalado docker y seguir los siguientes pasos

1. Crea el archivo .env y remplaza los valores por los deseados

### Nota: Los comandos se deben ejecutar en la carpeta raiz del proyecto
2. Construye la imagen:
`docker compose build`

3. Monta el contenedor:
`docker compose up -d`

4. Para detener el contenedor:
`docker compose down`


## Queries de ejemplo
Estos son los queries que pueden ser ejecutados para testear los diferentes endpoints existentes

1. Creacion de producto:
```
mutation {
  createProduct(
    input: {
      name: "Keyboard"
      price: 19.99
      stock: 40
    }
  ) {
    success
    message
  }
}
```

2. Obtencion de todos los productos:
```
query {
  products {
    id
    name
    price
    stock
    createdAt
  }
}
```

3. Obtencion de un producto por ID:
```
query {
  product(
    id: "1"
  ) {
    id
    name
    price
    stock
    createdAt
  }
}
```

4. Actualizacion de un producto:
```
mutation {
  updateProduct(
    id: "1"
    input: {
      name: "Mou"
    }
  ) {
    success
    message
  }
}
```

5. Eliminacion de un producto:
```
mutation{
  deleteProduct(id: "1"){
    success
    message
  }
}
```

Realizado por Edson Jair Fuentes García
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![CMake](https://img.shields.io/badge/CMake-%23008FBA.svg?style=for-the-badge&logo=cmake&logoColor=white)
![Shell Script](https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white)

# Backend Challenge Online-Store-API

__RestAPI written in Go. see documentation [Online-Store-API](http://ec2-3-94-76-231.compute-1.amazonaws.com) for details.__

Endpoint | Description | Progress |
--------|------------|--------|
[/api/v1/register](http://ec2-3-94-76-231.compute-1.amazonaws.com/#/auth/Register) | Customer can Register |<ul><li>- [x] Completed</li></ul>|
[/api/v1/login](http://ec2-3-94-76-231.compute-1.amazonaws.com/#/auth/Login) | Customer can Login |<ul><li>- [x] Completed</li></ul>|
[/api/v1/product](http://ec2-3-94-76-231.compute-1.amazonaws.com/#/product/ListProduct) | Customer can view product list by product category |<ul><li>- [x] Completed</li></ul>|
[/api/v1/cart](http://ec2-3-94-76-231.compute-1.amazonaws.com/#/cart/AddToCart) | Customer can add product to shopping cart |<ul><li>- [x] Completed</li></ul>| 
[/api/v1/cart](http://ec2-3-94-76-231.compute-1.amazonaws.com/#/cart/GetCart) | Customers can see a list of products that have been added to the shopping cart |<ul><li>- [x] Completed</li></ul>|
[/api/v1/cart](http://ec2-3-94-76-231.compute-1.amazonaws.com/#/cart/DeleteProductFromCart) | Customer can delete product list in shopping cart |<ul><li>- [x] Completed</li></ul>|
[/api/v1/order](http://ec2-3-94-76-231.compute-1.amazonaws.com/#/order/AddOrder) | Customers can checkout |<ul><li>- [x] Completed</li></ul>|
[/api/v1/order/proof](http://ec2-3-94-76-231.compute-1.amazonaws.com/#/order/AddProofPayment) | Customers can make payment transactions |<ul><li>- [x] Completed</li></ul>| 

## Knowledge
* [Go-lang](https://go.dev/doc/)
* [AWS EC2](https://aws.amazon.com/id/ec2/?trk=dee41712-cf2c-4ae2-9ca5-f793b9519b91&sc_channel=ps&ef_id=CjwKCAiA6byqBhAWEiwAnGCA4BwWB6f48cCUsURXAIZAQwhzJSA1lcgxxEVGi_SzwRzgvOg3gWamnRoCOwEQAvD_BwE:G:s&s_kwcid=AL!4422!3!590023635259!p!!g!!server%20amazon%20ec2!16178326774!136912369847)
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://www.docker.com/)
* [SSH](https://www.openssh.com/)
* [Github-Actions CI/CD](https://docs.github.com/en/actions)
* [golang-migrate](https://github.com/golang-migrate/migrate)
* [sqlc](https://github.com/sqlc-dev/sqlc)
* [oapi-codegen](https://github.com/deepmap/oapi-codegen)
* [gin](https://github.com/gin-gonic/gin)

# Get Started
Before starting, you must provide a VPS with minimum specifications:<br/>
`200MB RAM with 1GB storage`. and a little knowledge about port settings

## Setup Repository
1. Fork respository `git clone git@github.com:gafar-code/online-store.git`
2. Copy SSH host to Github Repository secret actions with name `SSH_HOST`
3. Copy SSH username to Github repository Secret Actions with name `SSH_USERNAME`<br/>

That's all you have to do for the Repository

### Setting App
You need to change a little configuration in the /doc/open-api.yaml file:<br/>
```ruby
servers:
  - url: http://ec2-3-94-76-231.compute-1.amazonaws.com:8080/api/v1/
```
Change the url in line 9 to your VPS url, for Example:
```ruby
servers:
  - url: http://your_vps_url:8080/api/v1/
```
This needs to be done so that swagger-ui can recognize your routing. That's all you need to change in the application. In fact, the application uses Docker Compose so that all settings have been implemented practically, let's go to the next step!

### Setup VPS
1. Make sure ports 8080 and 5432 are open on your VPS
2. Install Docker and Docker Compose to your VPS, See documentation [Docker](https://docs.docker.com/engine/install/ubuntu)
3. Set docker run on startup `sudo systemctl enable docker`
4. Create SSH key `ssh-keygen -t rsa -b 4096 -C "your_email@example.com"`
5. Copy SSH private key to Github Repository secret actions with name `SSH_PRIVATE_KEY`
6. Copy SSH public key to Github SSH key, see more [How to add New SSH key to Github](https://docs.github.com/en/github-ae@latest/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)
7. Clone repository using SSH `git clone git@github.com:your_username/online-store.git`
8. Goto root app directory `cd online-store`
9. Build and run `docker compose up`

Yeay! Your server is now ready to use!

## How to access RDS (Remote Database)
This is very easy to do. It's up to you if you want to use DBeaver, TablePlus or other software, it can still work well. The following is the default RDS configuration that you can use
```
HOST:     your_vps_address
PORT:     5432
USERNAME: root
PASSWORD: secret
DATABASE: store_db
```
<br/><br/><br/><br/><br/><br/>

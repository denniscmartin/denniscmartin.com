TITLE=Parsing financial reports
DESCRIPTION=Parsing financial reports
DATE=30/04/2023
+++

Mi intento de automatizar la extracción de los datos de un balance financiero en PDF a un formato estructurado, ponerlos a disposición de los usuarios a través de una API, y mostrarlos en una aplicación web.

He desarrollado la infraestructura en AWS utilizando los siguientes servicios: AWS Lambda, AWS Api Gateway, AWS Step Functions, AWS Textract, AWS S3, AWS Eventbridge, y AWS DynamoDB. El backend ha sido programado en Python y el frontend en Vuejs.

Code: <https://github.com/denniscmartin/finance-parser>

Youtube video: <https://www.youtube.com/watch?v=iBAKh0oA0ew>

![Frontend](./images/parsing-financial-reports.png)
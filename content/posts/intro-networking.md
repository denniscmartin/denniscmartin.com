+++
title = 'Introducción a las redes de ordenadores'
date = 2023-09-23
thumbnail = 'posts-intro-networking-thumbnail.png'
draft = false
tags = ['networking', 'internet']
+++

## Contexto

Hoy en día parece obvio que ordenadores de distintas marcas se puedan comunicar entre ellos. Pero no siempre ha sido así. En 1974, IBM publicó un modelo de comunicación llamado Arquitectura de Redes de Sistemas o (SNA) por sus siglas en inglés. Este modelo detallaba una serie de protocolos y especificaciones que permitían la comunicación entre ordenadores de IBM. El problema era que si no tenías un ordenador IBM no podías usar este modelo.

En los años siguientes, los demás fabricantes de ordenadores siguieron la misma lógica de IBM, creando modelos de comunicación cerrados que solo podían utilizar los ordenadores de la propia marca. Al principio estas redes cerradas funcionaron bien, pero después incrementaría la complejidad de las comunicaciones. Si tu empresa compraba ordenadores de tres marcas distintas, los ingenieros tenían que crear tres redes diferentes y luego intentar conectar esas tres redes entre ellas.

La solución sería un nuevo modelo abierto e independiente de cualquier empresa. A finales de los años 70, la Organización Internacional de la Normalización (ISO) empezó a desarrollar este modelo, el cual se convertiría en el modelo de Interconexión de Sistemas Abiertos, más conocido como (OSI). El objetivo de este modelo era claro. Estandarizar los protocolos de comunicación y así permitir la comunicación entre todos los ordenadores del mundo. ISO contó con participantes de los países más avanzados tecnológicamente para llevar a cabo esta tarea.

Mientras tanto, un segundo modelo abierto se empezaba a desarrollar a través de un contrato con el Departamento de Defensa de los Estados Unidos. Investigadores de varias universidades trabajaron voluntariamente en el desarrollo de estos protocolos, lo que resultó en la creación del modelo llamado TCP/IP.

Durante la década de los 90, las compañías empezaron a incluir los modelos OSI y TCP/IP, a sus redes empresariales. Sin embargo, a finales de los 90, TCP/IP se convirtió en el predominante y OSI empezaba a desaparecer. Volviendo al presente, TCP/IP es claramente el ganador. Los modelos cerrados aún existen, pero han sido mayormente reemplazados. El modelo OSI sufrió de un proceso de formalización muy lento comparado con TCP/IP, así que nunca tuvo éxito en el mercado. Al final, el modelo desarrollado por un puñado de voluntarios se ha convertido en el modelo de comunicación de ordenadores más importante hasta la fecha.

## TCP/IP

Este modelo define una serie de reglas que nos permiten conectar nuestro ordenador a la red. Podemos abrir el navegador que utilicemos y visitar cualquier página web. Simplemente, funciona. Esto es gracias a toda una pila de tecnología que ha sido desarrollada en base a los protocolos del modelo TCP/IP. Por ejemplo, los desarrolladores del sistema operativo que usas han implementado las partes necesarias de este modelo para que puedas comunicarte con otros ordenadores. A su vez, los fabricantes de las tarjetas de red han tenido en cuenta los estándares LAN referenciados por este modelo al fabricar sus productos. En resumen, los fabricantes y desarrolladores han implementado el modelo TCP/IP en sus productos.

Para simplificar el modelo, este se divide en varias capas. Cada capa incluye protocolos y estándares relaciones con una funcionalidad. El modelo TCP/IP tiene 5 capas, empezando por la física, enlace de datos, red, transporte, y aplicación. Normalmente se representa en una torre con una capa encima de la otra, de forma jerárquica, y se añade el número de cada capa.

### Capa de aplicación.

Empezando por la última, esta capa actúa como una interfaz entre el software que se ejecuta en tu ordenador y la red. Por ejemplo, dentro de esta capa está el protocolo HTTP, que define como los navegadores pueden solicitar contenido a un servidor web.

### Capa de transporte.

Si vamos a la capa anterior a la de aplicación, nos encontramos la capa de transporte. Los dos protocolos de transporte más utilizados son el Protocolo de Transmisión de Control (TCP) y el Protocolo de Datagramas de Usuario (UDP). Como todas las capas en este modelo, la capa de transporte proporciona servicios a la capa que está encima. Un ejemplo de estos servicios es la recuperación de errores en el protocolo TCP.

Imagínate que entras a una página web a través de tu navegador, pero el mensaje de respuesta del servidor se pierde por el camino. Si no hay ningún mecanismo para detectar y corregir los errores, tu navegador se quedaría esperando a recibir el mensaje para siempre. TCP proporciona estos mecanismos para que el protocolo HTTP los utilice y tu navegador no se quede esperando para siempre. Esta forma de interacción entre capas se conoce como interacción con la capa adyacente.

### Capa de red.

En la capa de red, probablemente el protocolo más conocido sea el Protocolo de Internet, o también llamado IP. El protocolo IP proporciona importantes características, las más importantes son direccionamiento y enrutamiento.

Por ejemplo, si nos queremos comunicar con otros ordenadores, necesitamos alguna forma de identificarlos. Por esta razón, a cada dispositivo conectado a la red se le asigna una dirección IP única. El proceso de enrutamiento empezaría con tu ordenador enviando a tu router tu mensaje junto con la dirección IP de destino, tu router comparará la dirección IP con una lista de direcciones ya conocidas. Probablemente tu router no conozca la IP entera pero sí parte de ella. Con esta información, tu router le enviará el mensaje a otro router. Este segundo router hará lo mismo. Seguirán comparando la IP hasta que un router conozca exactamente la dirección IP de destino, entonces el mensaje habrá llegado correctamente.

En la realidad es algo más complejo que esto, y normalmente hay otro tipo de dispositivos en la red, no solo routers, pero la idea es esa.

### Capa física y de enlace de datos.

Aunque son dos capas distintas, están fuertemente relacionadas. La capa física define el tipo de cableado, voltajes, y demás especificaciones de las señales eléctricas que fluyen por los cables. La capa de enlace de datos se centra en enviar los datos por el medio adecuado. Por ejemplo, una red por cable (Ethernet LAN) utiliza distintos protocolos de enlace de datos que una red inalámbrica (WIFI).

## Final

Pero ¿qué fue del modelo OSI? Pues bien, estas son las capas del modelo OSI. Al igual que TCP/IP, cada capa proporciona servicios a la capa que tiene encima. Además, comparte otras similitudes con TCP/IP. Sin embargo, la primera diferencia visible es el número de capas. TCP/IP tiene 5 capas, y OSI 7. Las 4 primeras son iguales, pero subdivide la capa de aplicación en tres. Sesión, presentación, y aplicación.

Gracias a esto, los ingenieros se refieren a la capa de aplicación en el modelo TCP/IP como la capa número 7, incluso cuando este modelo solo tiene 5 capas. ¿Por qué? Para confundir a la gente supongo. Tan listos para unas cosas y tan tontos para otras.

Youtube video: <https://www.youtube.com/watch?v=BkSUMRKRm5M>

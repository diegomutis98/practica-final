const botonBuscarEstudiante = document.querySelector(".search button");
const Id = document.querySelector("input");
const nombre = document.querySelector(".nombreRes");
const identidad = document.querySelector(".identidadRes");
const programa = document.querySelector(".programaRes");
const semestre = document.querySelector(".semestreRes");
const situacion = document.querySelector(".situacionRes");
const creditos = document.querySelector(".creditosRes");
const nivel = document.querySelector(".nivelRes");
const botonCrearEstudiante = document.querySelector(".crear button");
const botonEliminarEstudiante = document.querySelector(".eliminar button");
const botoneModificarEstudiante = document.querySelector(".modificar button");

// Para realisar la solicitid tipo Get
botonBuscarEstudiante.addEventListener("click", async () => {
    try {
    // Realiza la solicitud a la API principal
    const response = await axios.get(
      `http://localhost:8080/estudiantes/${Id.value}`
          );
    // leer datos de la BD
    nombre.innerHTML = response.data["nombre"];
    identidad.innerHTML = response.data["identidad"];
    programa.innerHTML = response.data["programa"];
    semestre.innerHTML = response.data["semestre"];
    situacion.innerHTML = response.data["situacion"];
    creditos.innerHTML = response.data["creditos"];
    nivel.innerHTML = response.data["nivel"];
    console.log(typeof(nombre.value))

    window.alert("Se cargarón correctamente los datos del estudiante ") 
  } catch (error) {
     // Manejo de errores   
     window.alert("Error al buscar el estudiante, es posible que ese Id no exista") 
  }
});

// Para realizar la solicitud de tipo Post
botonCrearEstudiante.addEventListener("click", async () => {
  try {
    const usuarioEnv = document.querySelector("#usuarioEnv" ).value
    const nombreEnv = document.querySelector("#nombreEnv" ).value
    const identidadEnv = document.querySelector("#identidadEnv" ).value
    const programaEnv = document.querySelector("#programaEnv" ).value
    const semestreEnv = Number(document.querySelector("#semestreEnv" ).value)
    const situacionEnv =document.querySelector("#situacionEnv" ).value
    const creditosEnv =document.querySelector("#creditosEnv" ).value
    const nivelEnv = document.querySelector("#nivelEnv" ).value
    
    const estudianteEnv = {usuario:`${usuarioEnv}`,
    nombre:`${nombreEnv}`,
    identidad:`${identidadEnv}`,
    programa:`${programaEnv}`,
    semestre: semestreEnv,
    situacion:`${situacionEnv}`,
    creditos:`${creditosEnv}`,
    nivel:`${nivelEnv}`,}

    console.log (estudianteEnv)
    console.log(typeof(usuarioEnv),typeof(nombreEnv),typeof(identidadEnv),
  typeof(programaEnv),typeof(semestreEnv),typeof(situacionEnv),typeof(creditosEnv),typeof(nivelEnv)
  )

  // Realiza la solicitud a la API principal
 const response = await axios.post(
    `http://localhost:8080/estudiantes`,estudianteEnv
        );
        window.alert("Los datos del estudiante fueron almacenados correctamente") 
} catch (error) {
   // Manejo de errores 
   window.alert("Error al crear al registar el nuevo estudiante, intente de nuevo por favor") 
}
});

// Para realizar la solicitud de tipo Delate
botonEliminarEstudiante.addEventListener("click", async () => {
  try {
  // Realiza la solicitud a la API principal
 const response = await axios.delete(
    `http://localhost:8080/estudiantes/${Id.value}`
        );
        console.log("si esta eliminando a un estudiante")
        window.alert("Se elimino al estudiante de la base de datos correctamente") 
  } catch (error) {
   // Manejo de errores  
   window.alert("Error al eliminar al estudiante, es posible que el id no se el correcto") 
}
});
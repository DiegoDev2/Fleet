package main

// MongodbAtlasCli - Atlas CLI enables you to manage your MongoDB Atlas
// Homepage: https://www.mongodb.com/docs/atlas/cli/stable/

import (
	"fmt"
	
	"os/exec"
)

func installMongodbAtlasCli() {
	// Método 1: Descargar y extraer .tar.gz
	mongodbatlascli_tar_url := "https://github.com/mongodb/mongodb-atlas-cli/archive/refs/tags/atlascli/v1.27.0.tar.gz"
	mongodbatlascli_cmd_tar := exec.Command("curl", "-L", mongodbatlascli_tar_url, "-o", "package.tar.gz")
	err := mongodbatlascli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongodbatlascli_zip_url := "https://github.com/mongodb/mongodb-atlas-cli/archive/refs/tags/atlascli/v1.27.0.zip"
	mongodbatlascli_cmd_zip := exec.Command("curl", "-L", mongodbatlascli_zip_url, "-o", "package.zip")
	err = mongodbatlascli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongodbatlascli_bin_url := "https://github.com/mongodb/mongodb-atlas-cli/archive/refs/tags/atlascli/v1.27.0.bin"
	mongodbatlascli_cmd_bin := exec.Command("curl", "-L", mongodbatlascli_bin_url, "-o", "binary.bin")
	err = mongodbatlascli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongodbatlascli_src_url := "https://github.com/mongodb/mongodb-atlas-cli/archive/refs/tags/atlascli/v1.27.0.src.tar.gz"
	mongodbatlascli_cmd_src := exec.Command("curl", "-L", mongodbatlascli_src_url, "-o", "source.tar.gz")
	err = mongodbatlascli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongodbatlascli_cmd_direct := exec.Command("./binary")
	err = mongodbatlascli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: mongosh")
exec.Command("latte", "install", "mongosh")
}

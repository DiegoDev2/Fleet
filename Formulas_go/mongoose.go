package main

// Mongoose - Web server build on top of Libmongoose embedded library
// Homepage: https://mongoose.ws/

import (
	"fmt"
	
	"os/exec"
)

func installMongoose() {
	// Método 1: Descargar y extraer .tar.gz
	mongoose_tar_url := "https://github.com/cesanta/mongoose/archive/refs/tags/7.15.tar.gz"
	mongoose_cmd_tar := exec.Command("curl", "-L", mongoose_tar_url, "-o", "package.tar.gz")
	err := mongoose_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongoose_zip_url := "https://github.com/cesanta/mongoose/archive/refs/tags/7.15.zip"
	mongoose_cmd_zip := exec.Command("curl", "-L", mongoose_zip_url, "-o", "package.zip")
	err = mongoose_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongoose_bin_url := "https://github.com/cesanta/mongoose/archive/refs/tags/7.15.bin"
	mongoose_cmd_bin := exec.Command("curl", "-L", mongoose_bin_url, "-o", "binary.bin")
	err = mongoose_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongoose_src_url := "https://github.com/cesanta/mongoose/archive/refs/tags/7.15.src.tar.gz"
	mongoose_cmd_src := exec.Command("curl", "-L", mongoose_src_url, "-o", "source.tar.gz")
	err = mongoose_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongoose_cmd_direct := exec.Command("./binary")
	err = mongoose_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}

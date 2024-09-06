package main

// MongoCDriver - C driver for MongoDB
// Homepage: https://github.com/mongodb/mongo-c-driver

import (
	"fmt"
	
	"os/exec"
)

func installMongoCDriver() {
	// Método 1: Descargar y extraer .tar.gz
	mongocdriver_tar_url := "https://github.com/mongodb/mongo-c-driver/archive/refs/tags/1.27.6.tar.gz"
	mongocdriver_cmd_tar := exec.Command("curl", "-L", mongocdriver_tar_url, "-o", "package.tar.gz")
	err := mongocdriver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongocdriver_zip_url := "https://github.com/mongodb/mongo-c-driver/archive/refs/tags/1.27.6.zip"
	mongocdriver_cmd_zip := exec.Command("curl", "-L", mongocdriver_zip_url, "-o", "package.zip")
	err = mongocdriver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongocdriver_bin_url := "https://github.com/mongodb/mongo-c-driver/archive/refs/tags/1.27.6.bin"
	mongocdriver_cmd_bin := exec.Command("curl", "-L", mongocdriver_bin_url, "-o", "binary.bin")
	err = mongocdriver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongocdriver_src_url := "https://github.com/mongodb/mongo-c-driver/archive/refs/tags/1.27.6.src.tar.gz"
	mongocdriver_cmd_src := exec.Command("curl", "-L", mongocdriver_src_url, "-o", "source.tar.gz")
	err = mongocdriver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongocdriver_cmd_direct := exec.Command("./binary")
	err = mongocdriver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}

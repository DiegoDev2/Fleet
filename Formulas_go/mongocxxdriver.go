package main

// MongoCxxDriver - C++ driver for MongoDB
// Homepage: https://github.com/mongodb/mongo-cxx-driver

import (
	"fmt"
	
	"os/exec"
)

func installMongoCxxDriver() {
	// Método 1: Descargar y extraer .tar.gz
	mongocxxdriver_tar_url := "https://github.com/mongodb/mongo-cxx-driver/releases/download/r3.10.2/mongo-cxx-driver-r3.10.2.tar.gz"
	mongocxxdriver_cmd_tar := exec.Command("curl", "-L", mongocxxdriver_tar_url, "-o", "package.tar.gz")
	err := mongocxxdriver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongocxxdriver_zip_url := "https://github.com/mongodb/mongo-cxx-driver/releases/download/r3.10.2/mongo-cxx-driver-r3.10.2.zip"
	mongocxxdriver_cmd_zip := exec.Command("curl", "-L", mongocxxdriver_zip_url, "-o", "package.zip")
	err = mongocxxdriver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongocxxdriver_bin_url := "https://github.com/mongodb/mongo-cxx-driver/releases/download/r3.10.2/mongo-cxx-driver-r3.10.2.bin"
	mongocxxdriver_cmd_bin := exec.Command("curl", "-L", mongocxxdriver_bin_url, "-o", "binary.bin")
	err = mongocxxdriver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongocxxdriver_src_url := "https://github.com/mongodb/mongo-cxx-driver/releases/download/r3.10.2/mongo-cxx-driver-r3.10.2.src.tar.gz"
	mongocxxdriver_cmd_src := exec.Command("curl", "-L", mongocxxdriver_src_url, "-o", "source.tar.gz")
	err = mongocxxdriver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongocxxdriver_cmd_direct := exec.Command("./binary")
	err = mongocxxdriver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: mongo-c-driver")
exec.Command("latte", "install", "mongo-c-driver")
}

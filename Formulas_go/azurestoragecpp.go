package main

// AzureStorageCpp - Microsoft Azure Storage Client Library for C++
// Homepage: https://azure.github.io/azure-storage-cpp

import (
	"fmt"
	
	"os/exec"
)

func installAzureStorageCpp() {
	// Método 1: Descargar y extraer .tar.gz
	azurestoragecpp_tar_url := "https://github.com/Azure/azure-storage-cpp/archive/refs/tags/v7.5.0.tar.gz"
	azurestoragecpp_cmd_tar := exec.Command("curl", "-L", azurestoragecpp_tar_url, "-o", "package.tar.gz")
	err := azurestoragecpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	azurestoragecpp_zip_url := "https://github.com/Azure/azure-storage-cpp/archive/refs/tags/v7.5.0.zip"
	azurestoragecpp_cmd_zip := exec.Command("curl", "-L", azurestoragecpp_zip_url, "-o", "package.zip")
	err = azurestoragecpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	azurestoragecpp_bin_url := "https://github.com/Azure/azure-storage-cpp/archive/refs/tags/v7.5.0.bin"
	azurestoragecpp_cmd_bin := exec.Command("curl", "-L", azurestoragecpp_bin_url, "-o", "binary.bin")
	err = azurestoragecpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	azurestoragecpp_src_url := "https://github.com/Azure/azure-storage-cpp/archive/refs/tags/v7.5.0.src.tar.gz"
	azurestoragecpp_cmd_src := exec.Command("curl", "-L", azurestoragecpp_src_url, "-o", "source.tar.gz")
	err = azurestoragecpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	azurestoragecpp_cmd_direct := exec.Command("./binary")
	err = azurestoragecpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cpprestsdk")
exec.Command("latte", "install", "cpprestsdk")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}

package main

// Yazpp - C++ API for the Yaz toolkit
// Homepage: https://www.indexdata.com/resources/software/yazpp/

import (
	"fmt"
	
	"os/exec"
)

func installYazpp() {
	// Método 1: Descargar y extraer .tar.gz
	yazpp_tar_url := "https://ftp.indexdata.com/pub/yazpp/yazpp-1.8.1.tar.gz"
	yazpp_cmd_tar := exec.Command("curl", "-L", yazpp_tar_url, "-o", "package.tar.gz")
	err := yazpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yazpp_zip_url := "https://ftp.indexdata.com/pub/yazpp/yazpp-1.8.1.zip"
	yazpp_cmd_zip := exec.Command("curl", "-L", yazpp_zip_url, "-o", "package.zip")
	err = yazpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yazpp_bin_url := "https://ftp.indexdata.com/pub/yazpp/yazpp-1.8.1.bin"
	yazpp_cmd_bin := exec.Command("curl", "-L", yazpp_bin_url, "-o", "binary.bin")
	err = yazpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yazpp_src_url := "https://ftp.indexdata.com/pub/yazpp/yazpp-1.8.1.src.tar.gz"
	yazpp_cmd_src := exec.Command("curl", "-L", yazpp_src_url, "-o", "source.tar.gz")
	err = yazpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yazpp_cmd_direct := exec.Command("./binary")
	err = yazpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: yaz")
	exec.Command("latte", "install", "yaz").Run()
}

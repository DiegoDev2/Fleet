package main

// Grantlee - Libraries for text templating with Qt
// Homepage: https://github.com/steveire/grantlee

import (
	"fmt"
	
	"os/exec"
)

func installGrantlee() {
	// Método 1: Descargar y extraer .tar.gz
	grantlee_tar_url := "https://github.com/steveire/grantlee/releases/download/v5.3.1/grantlee-5.3.1.tar.gz"
	grantlee_cmd_tar := exec.Command("curl", "-L", grantlee_tar_url, "-o", "package.tar.gz")
	err := grantlee_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grantlee_zip_url := "https://github.com/steveire/grantlee/releases/download/v5.3.1/grantlee-5.3.1.zip"
	grantlee_cmd_zip := exec.Command("curl", "-L", grantlee_zip_url, "-o", "package.zip")
	err = grantlee_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grantlee_bin_url := "https://github.com/steveire/grantlee/releases/download/v5.3.1/grantlee-5.3.1.bin"
	grantlee_cmd_bin := exec.Command("curl", "-L", grantlee_bin_url, "-o", "binary.bin")
	err = grantlee_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grantlee_src_url := "https://github.com/steveire/grantlee/releases/download/v5.3.1/grantlee-5.3.1.src.tar.gz"
	grantlee_cmd_src := exec.Command("curl", "-L", grantlee_src_url, "-o", "source.tar.gz")
	err = grantlee_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grantlee_cmd_direct := exec.Command("./binary")
	err = grantlee_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: graphviz")
	exec.Command("latte", "install", "graphviz").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
}

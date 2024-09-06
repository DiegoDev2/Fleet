package main

// Lasso - Library for Liberty Alliance and SAML protocols
// Homepage: https://lasso.entrouvert.org/

import (
	"fmt"
	
	"os/exec"
)

func installLasso() {
	// Método 1: Descargar y extraer .tar.gz
	lasso_tar_url := "https://dev.entrouvert.org/releases/lasso/lasso-2.8.2.tar.gz"
	lasso_cmd_tar := exec.Command("curl", "-L", lasso_tar_url, "-o", "package.tar.gz")
	err := lasso_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lasso_zip_url := "https://dev.entrouvert.org/releases/lasso/lasso-2.8.2.zip"
	lasso_cmd_zip := exec.Command("curl", "-L", lasso_zip_url, "-o", "package.zip")
	err = lasso_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lasso_bin_url := "https://dev.entrouvert.org/releases/lasso/lasso-2.8.2.bin"
	lasso_cmd_bin := exec.Command("curl", "-L", lasso_bin_url, "-o", "binary.bin")
	err = lasso_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lasso_src_url := "https://dev.entrouvert.org/releases/lasso/lasso-2.8.2.src.tar.gz"
	lasso_cmd_src := exec.Command("curl", "-L", lasso_src_url, "-o", "source.tar.gz")
	err = lasso_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lasso_cmd_direct := exec.Command("./binary")
	err = lasso_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: libxmlsec1")
	exec.Command("latte", "install", "libxmlsec1").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}

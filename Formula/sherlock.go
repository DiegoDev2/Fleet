package main

// Sherlock - Hunt down social media accounts by username
// Homepage: https://sherlockproject.xyz/

import (
	"fmt"
	
	"os/exec"
)

func installSherlock() {
	// Método 1: Descargar y extraer .tar.gz
	sherlock_tar_url := "https://files.pythonhosted.org/packages/0a/95/b4f7a399c43d1d57a703ddf08513411bbb0bfc6bbaabab7ad4e2c534bba7/sherlock_project-0.15.0.tar.gz"
	sherlock_cmd_tar := exec.Command("curl", "-L", sherlock_tar_url, "-o", "package.tar.gz")
	err := sherlock_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sherlock_zip_url := "https://files.pythonhosted.org/packages/0a/95/b4f7a399c43d1d57a703ddf08513411bbb0bfc6bbaabab7ad4e2c534bba7/sherlock_project-0.15.0.zip"
	sherlock_cmd_zip := exec.Command("curl", "-L", sherlock_zip_url, "-o", "package.zip")
	err = sherlock_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sherlock_bin_url := "https://files.pythonhosted.org/packages/0a/95/b4f7a399c43d1d57a703ddf08513411bbb0bfc6bbaabab7ad4e2c534bba7/sherlock_project-0.15.0.bin"
	sherlock_cmd_bin := exec.Command("curl", "-L", sherlock_bin_url, "-o", "binary.bin")
	err = sherlock_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sherlock_src_url := "https://files.pythonhosted.org/packages/0a/95/b4f7a399c43d1d57a703ddf08513411bbb0bfc6bbaabab7ad4e2c534bba7/sherlock_project-0.15.0.src.tar.gz"
	sherlock_cmd_src := exec.Command("curl", "-L", sherlock_src_url, "-o", "source.tar.gz")
	err = sherlock_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sherlock_cmd_direct := exec.Command("./binary")
	err = sherlock_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}

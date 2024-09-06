package main

// Page - Use Neovim as pager
// Homepage: https://github.com/I60R/page

import (
	"fmt"
	
	"os/exec"
)

func installPage() {
	// Método 1: Descargar y extraer .tar.gz
	page_tar_url := "https://github.com/I60R/page/archive/refs/tags/v4.6.3.tar.gz"
	page_cmd_tar := exec.Command("curl", "-L", page_tar_url, "-o", "package.tar.gz")
	err := page_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	page_zip_url := "https://github.com/I60R/page/archive/refs/tags/v4.6.3.zip"
	page_cmd_zip := exec.Command("curl", "-L", page_zip_url, "-o", "package.zip")
	err = page_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	page_bin_url := "https://github.com/I60R/page/archive/refs/tags/v4.6.3.bin"
	page_cmd_bin := exec.Command("curl", "-L", page_bin_url, "-o", "binary.bin")
	err = page_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	page_src_url := "https://github.com/I60R/page/archive/refs/tags/v4.6.3.src.tar.gz"
	page_cmd_src := exec.Command("curl", "-L", page_src_url, "-o", "source.tar.gz")
	err = page_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	page_cmd_direct := exec.Command("./binary")
	err = page_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: neovim")
	exec.Command("latte", "install", "neovim").Run()
}

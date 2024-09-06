package main

// Blogc - Blog compiler with template engine and markup language
// Homepage: https://blogc.rgm.io/

import (
	"fmt"
	
	"os/exec"
)

func installBlogc() {
	// Método 1: Descargar y extraer .tar.gz
	blogc_tar_url := "https://github.com/blogc/blogc/releases/download/v0.20.1/blogc-0.20.1.tar.xz"
	blogc_cmd_tar := exec.Command("curl", "-L", blogc_tar_url, "-o", "package.tar.gz")
	err := blogc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blogc_zip_url := "https://github.com/blogc/blogc/releases/download/v0.20.1/blogc-0.20.1.tar.xz"
	blogc_cmd_zip := exec.Command("curl", "-L", blogc_zip_url, "-o", "package.zip")
	err = blogc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blogc_bin_url := "https://github.com/blogc/blogc/releases/download/v0.20.1/blogc-0.20.1.tar.xz"
	blogc_cmd_bin := exec.Command("curl", "-L", blogc_bin_url, "-o", "binary.bin")
	err = blogc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blogc_src_url := "https://github.com/blogc/blogc/releases/download/v0.20.1/blogc-0.20.1.tar.xz"
	blogc_cmd_src := exec.Command("curl", "-L", blogc_src_url, "-o", "source.tar.gz")
	err = blogc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blogc_cmd_direct := exec.Command("./binary")
	err = blogc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

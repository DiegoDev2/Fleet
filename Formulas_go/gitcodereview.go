package main

// GitCodereview - Tool for working with Gerrit code reviews
// Homepage: https://pkg.go.dev/golang.org/x/review/git-codereview

import (
	"fmt"
	
	"os/exec"
)

func installGitCodereview() {
	// Método 1: Descargar y extraer .tar.gz
	gitcodereview_tar_url := "https://github.com/golang/review/archive/refs/tags/v1.12.0.tar.gz"
	gitcodereview_cmd_tar := exec.Command("curl", "-L", gitcodereview_tar_url, "-o", "package.tar.gz")
	err := gitcodereview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitcodereview_zip_url := "https://github.com/golang/review/archive/refs/tags/v1.12.0.zip"
	gitcodereview_cmd_zip := exec.Command("curl", "-L", gitcodereview_zip_url, "-o", "package.zip")
	err = gitcodereview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitcodereview_bin_url := "https://github.com/golang/review/archive/refs/tags/v1.12.0.bin"
	gitcodereview_cmd_bin := exec.Command("curl", "-L", gitcodereview_bin_url, "-o", "binary.bin")
	err = gitcodereview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitcodereview_src_url := "https://github.com/golang/review/archive/refs/tags/v1.12.0.src.tar.gz"
	gitcodereview_cmd_src := exec.Command("curl", "-L", gitcodereview_src_url, "-o", "source.tar.gz")
	err = gitcodereview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitcodereview_cmd_direct := exec.Command("./binary")
	err = gitcodereview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}

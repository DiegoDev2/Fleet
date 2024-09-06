package main

// DjangoCompletion - Bash completion for Django
// Homepage: https://www.djangoproject.com/

import (
	"fmt"
	
	"os/exec"
)

func installDjangoCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	djangocompletion_tar_url := "https://github.com/django/django/archive/refs/tags/5.1.1.tar.gz"
	djangocompletion_cmd_tar := exec.Command("curl", "-L", djangocompletion_tar_url, "-o", "package.tar.gz")
	err := djangocompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	djangocompletion_zip_url := "https://github.com/django/django/archive/refs/tags/5.1.1.zip"
	djangocompletion_cmd_zip := exec.Command("curl", "-L", djangocompletion_zip_url, "-o", "package.zip")
	err = djangocompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	djangocompletion_bin_url := "https://github.com/django/django/archive/refs/tags/5.1.1.bin"
	djangocompletion_cmd_bin := exec.Command("curl", "-L", djangocompletion_bin_url, "-o", "binary.bin")
	err = djangocompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	djangocompletion_src_url := "https://github.com/django/django/archive/refs/tags/5.1.1.src.tar.gz"
	djangocompletion_cmd_src := exec.Command("curl", "-L", djangocompletion_src_url, "-o", "source.tar.gz")
	err = djangocompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	djangocompletion_cmd_direct := exec.Command("./binary")
	err = djangocompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

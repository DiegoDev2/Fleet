package main

// Plenv - Perl binary manager
// Homepage: https://github.com/tokuhirom/plenv

import (
	"fmt"
	
	"os/exec"
)

func installPlenv() {
	// Método 1: Descargar y extraer .tar.gz
	plenv_tar_url := "https://github.com/tokuhirom/plenv/archive/refs/tags/2.3.1.tar.gz"
	plenv_cmd_tar := exec.Command("curl", "-L", plenv_tar_url, "-o", "package.tar.gz")
	err := plenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plenv_zip_url := "https://github.com/tokuhirom/plenv/archive/refs/tags/2.3.1.zip"
	plenv_cmd_zip := exec.Command("curl", "-L", plenv_zip_url, "-o", "package.zip")
	err = plenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plenv_bin_url := "https://github.com/tokuhirom/plenv/archive/refs/tags/2.3.1.bin"
	plenv_cmd_bin := exec.Command("curl", "-L", plenv_bin_url, "-o", "binary.bin")
	err = plenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plenv_src_url := "https://github.com/tokuhirom/plenv/archive/refs/tags/2.3.1.src.tar.gz"
	plenv_cmd_src := exec.Command("curl", "-L", plenv_src_url, "-o", "source.tar.gz")
	err = plenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plenv_cmd_direct := exec.Command("./binary")
	err = plenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: perl-build")
exec.Command("latte", "install", "perl-build")
}

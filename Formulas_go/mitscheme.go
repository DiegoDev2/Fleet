package main

// MitScheme - MIT/GNU Scheme development tools and runtime library
// Homepage: https://www.gnu.org/software/mit-scheme/

import (
	"fmt"
	
	"os/exec"
)

func installMitScheme() {
	// Método 1: Descargar y extraer .tar.gz
	mitscheme_tar_url := "https://ftp.gnu.org/gnu/mit-scheme/stable.pkg/12.1/mit-scheme-12.1-svm1-64le.tar.gz"
	mitscheme_cmd_tar := exec.Command("curl", "-L", mitscheme_tar_url, "-o", "package.tar.gz")
	err := mitscheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mitscheme_zip_url := "https://ftp.gnu.org/gnu/mit-scheme/stable.pkg/12.1/mit-scheme-12.1-svm1-64le.zip"
	mitscheme_cmd_zip := exec.Command("curl", "-L", mitscheme_zip_url, "-o", "package.zip")
	err = mitscheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mitscheme_bin_url := "https://ftp.gnu.org/gnu/mit-scheme/stable.pkg/12.1/mit-scheme-12.1-svm1-64le.bin"
	mitscheme_cmd_bin := exec.Command("curl", "-L", mitscheme_bin_url, "-o", "binary.bin")
	err = mitscheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mitscheme_src_url := "https://ftp.gnu.org/gnu/mit-scheme/stable.pkg/12.1/mit-scheme-12.1-svm1-64le.src.tar.gz"
	mitscheme_cmd_src := exec.Command("curl", "-L", mitscheme_src_url, "-o", "source.tar.gz")
	err = mitscheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mitscheme_cmd_direct := exec.Command("./binary")
	err = mitscheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}

package main

// Enscript - Convert text to Postscript, HTML, or RTF, with syntax highlighting
// Homepage: https://www.gnu.org/software/enscript/

import (
	"fmt"
	
	"os/exec"
)

func installEnscript() {
	// Método 1: Descargar y extraer .tar.gz
	enscript_tar_url := "https://ftp.gnu.org/gnu/enscript/enscript-1.6.6.tar.gz"
	enscript_cmd_tar := exec.Command("curl", "-L", enscript_tar_url, "-o", "package.tar.gz")
	err := enscript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enscript_zip_url := "https://ftp.gnu.org/gnu/enscript/enscript-1.6.6.zip"
	enscript_cmd_zip := exec.Command("curl", "-L", enscript_zip_url, "-o", "package.zip")
	err = enscript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enscript_bin_url := "https://ftp.gnu.org/gnu/enscript/enscript-1.6.6.bin"
	enscript_cmd_bin := exec.Command("curl", "-L", enscript_bin_url, "-o", "binary.bin")
	err = enscript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enscript_src_url := "https://ftp.gnu.org/gnu/enscript/enscript-1.6.6.src.tar.gz"
	enscript_cmd_src := exec.Command("curl", "-L", enscript_src_url, "-o", "source.tar.gz")
	err = enscript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enscript_cmd_direct := exec.Command("./binary")
	err = enscript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}

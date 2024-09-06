package main

// Gettext - GNU internationalization (i18n) and localization (l10n) library
// Homepage: https://www.gnu.org/software/gettext/

import (
	"fmt"
	
	"os/exec"
)

func installGettext() {
	// Método 1: Descargar y extraer .tar.gz
	gettext_tar_url := "https://ftp.gnu.org/gnu/gettext/gettext-0.22.5.tar.gz"
	gettext_cmd_tar := exec.Command("curl", "-L", gettext_tar_url, "-o", "package.tar.gz")
	err := gettext_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gettext_zip_url := "https://ftp.gnu.org/gnu/gettext/gettext-0.22.5.zip"
	gettext_cmd_zip := exec.Command("curl", "-L", gettext_zip_url, "-o", "package.zip")
	err = gettext_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gettext_bin_url := "https://ftp.gnu.org/gnu/gettext/gettext-0.22.5.bin"
	gettext_cmd_bin := exec.Command("curl", "-L", gettext_bin_url, "-o", "binary.bin")
	err = gettext_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gettext_src_url := "https://ftp.gnu.org/gnu/gettext/gettext-0.22.5.src.tar.gz"
	gettext_cmd_src := exec.Command("curl", "-L", gettext_src_url, "-o", "source.tar.gz")
	err = gettext_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gettext_cmd_direct := exec.Command("./binary")
	err = gettext_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

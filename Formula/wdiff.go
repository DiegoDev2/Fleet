package main

// Wdiff - Display word differences between text files
// Homepage: https://www.gnu.org/software/wdiff/

import (
	"fmt"
	
	"os/exec"
)

func installWdiff() {
	// Método 1: Descargar y extraer .tar.gz
	wdiff_tar_url := "https://ftp.gnu.org/gnu/wdiff/wdiff-1.2.2.tar.gz"
	wdiff_cmd_tar := exec.Command("curl", "-L", wdiff_tar_url, "-o", "package.tar.gz")
	err := wdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wdiff_zip_url := "https://ftp.gnu.org/gnu/wdiff/wdiff-1.2.2.zip"
	wdiff_cmd_zip := exec.Command("curl", "-L", wdiff_zip_url, "-o", "package.zip")
	err = wdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wdiff_bin_url := "https://ftp.gnu.org/gnu/wdiff/wdiff-1.2.2.bin"
	wdiff_cmd_bin := exec.Command("curl", "-L", wdiff_bin_url, "-o", "binary.bin")
	err = wdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wdiff_src_url := "https://ftp.gnu.org/gnu/wdiff/wdiff-1.2.2.src.tar.gz"
	wdiff_cmd_src := exec.Command("curl", "-L", wdiff_src_url, "-o", "source.tar.gz")
	err = wdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wdiff_cmd_direct := exec.Command("./binary")
	err = wdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}

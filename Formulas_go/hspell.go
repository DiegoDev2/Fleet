package main

// Hspell - Free Hebrew linguistic project
// Homepage: http://hspell.ivrix.org.il/

import (
	"fmt"
	
	"os/exec"
)

func installHspell() {
	// Método 1: Descargar y extraer .tar.gz
	hspell_tar_url := "http://hspell.ivrix.org.il/hspell-1.4.tar.gz"
	hspell_cmd_tar := exec.Command("curl", "-L", hspell_tar_url, "-o", "package.tar.gz")
	err := hspell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hspell_zip_url := "http://hspell.ivrix.org.il/hspell-1.4.zip"
	hspell_cmd_zip := exec.Command("curl", "-L", hspell_zip_url, "-o", "package.zip")
	err = hspell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hspell_bin_url := "http://hspell.ivrix.org.il/hspell-1.4.bin"
	hspell_cmd_bin := exec.Command("curl", "-L", hspell_bin_url, "-o", "binary.bin")
	err = hspell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hspell_src_url := "http://hspell.ivrix.org.il/hspell-1.4.src.tar.gz"
	hspell_cmd_src := exec.Command("curl", "-L", hspell_src_url, "-o", "source.tar.gz")
	err = hspell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hspell_cmd_direct := exec.Command("./binary")
	err = hspell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
}

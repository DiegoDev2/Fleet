package main

// Shakespeare - Write programs in Shakespearean English
// Homepage: https://web.archive.org/web/20211106102807/https://sourceforge.net/projects/shakespearelang/

import (
	"fmt"
	
	"os/exec"
)

func installShakespeare() {
	// Método 1: Descargar y extraer .tar.gz
	shakespeare_tar_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/shakespeare/spl-1.2.1.tar.gz"
	shakespeare_cmd_tar := exec.Command("curl", "-L", shakespeare_tar_url, "-o", "package.tar.gz")
	err := shakespeare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shakespeare_zip_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/shakespeare/spl-1.2.1.zip"
	shakespeare_cmd_zip := exec.Command("curl", "-L", shakespeare_zip_url, "-o", "package.zip")
	err = shakespeare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shakespeare_bin_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/shakespeare/spl-1.2.1.bin"
	shakespeare_cmd_bin := exec.Command("curl", "-L", shakespeare_bin_url, "-o", "binary.bin")
	err = shakespeare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shakespeare_src_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/shakespeare/spl-1.2.1.src.tar.gz"
	shakespeare_cmd_src := exec.Command("curl", "-L", shakespeare_src_url, "-o", "source.tar.gz")
	err = shakespeare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shakespeare_cmd_direct := exec.Command("./binary")
	err = shakespeare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
}

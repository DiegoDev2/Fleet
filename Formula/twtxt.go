package main

// Twtxt - Decentralised, minimalist microblogging service for hackers
// Homepage: https://github.com/buckket/twtxt

import (
	"fmt"
	
	"os/exec"
)

func installTwtxt() {
	// Método 1: Descargar y extraer .tar.gz
	twtxt_tar_url := "https://files.pythonhosted.org/packages/fc/4c/cff74642212dbca8d4d9059119555cd335324b3da0b52990a414a0257756/twtxt-1.3.1.tar.gz"
	twtxt_cmd_tar := exec.Command("curl", "-L", twtxt_tar_url, "-o", "package.tar.gz")
	err := twtxt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	twtxt_zip_url := "https://files.pythonhosted.org/packages/fc/4c/cff74642212dbca8d4d9059119555cd335324b3da0b52990a414a0257756/twtxt-1.3.1.zip"
	twtxt_cmd_zip := exec.Command("curl", "-L", twtxt_zip_url, "-o", "package.zip")
	err = twtxt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	twtxt_bin_url := "https://files.pythonhosted.org/packages/fc/4c/cff74642212dbca8d4d9059119555cd335324b3da0b52990a414a0257756/twtxt-1.3.1.bin"
	twtxt_cmd_bin := exec.Command("curl", "-L", twtxt_bin_url, "-o", "binary.bin")
	err = twtxt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	twtxt_src_url := "https://files.pythonhosted.org/packages/fc/4c/cff74642212dbca8d4d9059119555cd335324b3da0b52990a414a0257756/twtxt-1.3.1.src.tar.gz"
	twtxt_cmd_src := exec.Command("curl", "-L", twtxt_src_url, "-o", "source.tar.gz")
	err = twtxt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	twtxt_cmd_direct := exec.Command("./binary")
	err = twtxt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}

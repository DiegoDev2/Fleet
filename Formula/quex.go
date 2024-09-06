package main

// Quex - Generate lexical analyzers
// Homepage: https://quex.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installQuex() {
	// Método 1: Descargar y extraer .tar.gz
	quex_tar_url := "https://downloads.sourceforge.net/project/quex/quex-0.71.2.zip"
	quex_cmd_tar := exec.Command("curl", "-L", quex_tar_url, "-o", "package.tar.gz")
	err := quex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quex_zip_url := "https://downloads.sourceforge.net/project/quex/quex-0.71.2.zip"
	quex_cmd_zip := exec.Command("curl", "-L", quex_zip_url, "-o", "package.zip")
	err = quex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quex_bin_url := "https://downloads.sourceforge.net/project/quex/quex-0.71.2.zip"
	quex_cmd_bin := exec.Command("curl", "-L", quex_bin_url, "-o", "binary.bin")
	err = quex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quex_src_url := "https://downloads.sourceforge.net/project/quex/quex-0.71.2.zip"
	quex_cmd_src := exec.Command("curl", "-L", quex_src_url, "-o", "source.tar.gz")
	err = quex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quex_cmd_direct := exec.Command("./binary")
	err = quex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}

package main

// Rtf2latex2e - RTF-to-LaTeX translation
// Homepage: https://rtf2latex2e.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installRtf2latex2e() {
	// Método 1: Descargar y extraer .tar.gz
	rtf2latex2e_tar_url := "https://downloads.sourceforge.net/project/rtf2latex2e/rtf2latex2e-unix/2-2/rtf2latex2e-2-2-3.tar.gz"
	rtf2latex2e_cmd_tar := exec.Command("curl", "-L", rtf2latex2e_tar_url, "-o", "package.tar.gz")
	err := rtf2latex2e_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rtf2latex2e_zip_url := "https://downloads.sourceforge.net/project/rtf2latex2e/rtf2latex2e-unix/2-2/rtf2latex2e-2-2-3.zip"
	rtf2latex2e_cmd_zip := exec.Command("curl", "-L", rtf2latex2e_zip_url, "-o", "package.zip")
	err = rtf2latex2e_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rtf2latex2e_bin_url := "https://downloads.sourceforge.net/project/rtf2latex2e/rtf2latex2e-unix/2-2/rtf2latex2e-2-2-3.bin"
	rtf2latex2e_cmd_bin := exec.Command("curl", "-L", rtf2latex2e_bin_url, "-o", "binary.bin")
	err = rtf2latex2e_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rtf2latex2e_src_url := "https://downloads.sourceforge.net/project/rtf2latex2e/rtf2latex2e-unix/2-2/rtf2latex2e-2-2-3.src.tar.gz"
	rtf2latex2e_cmd_src := exec.Command("curl", "-L", rtf2latex2e_src_url, "-o", "source.tar.gz")
	err = rtf2latex2e_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rtf2latex2e_cmd_direct := exec.Command("./binary")
	err = rtf2latex2e_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

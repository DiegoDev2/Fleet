package main

// Latex2rtf - Translate LaTeX to RTF
// Homepage: https://latex2rtf.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLatex2rtf() {
	// Método 1: Descargar y extraer .tar.gz
	latex2rtf_tar_url := "https://downloads.sourceforge.net/project/latex2rtf/latex2rtf-unix/2.3.18/latex2rtf-2.3.18a.tar.gz"
	latex2rtf_cmd_tar := exec.Command("curl", "-L", latex2rtf_tar_url, "-o", "package.tar.gz")
	err := latex2rtf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	latex2rtf_zip_url := "https://downloads.sourceforge.net/project/latex2rtf/latex2rtf-unix/2.3.18/latex2rtf-2.3.18a.zip"
	latex2rtf_cmd_zip := exec.Command("curl", "-L", latex2rtf_zip_url, "-o", "package.zip")
	err = latex2rtf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	latex2rtf_bin_url := "https://downloads.sourceforge.net/project/latex2rtf/latex2rtf-unix/2.3.18/latex2rtf-2.3.18a.bin"
	latex2rtf_cmd_bin := exec.Command("curl", "-L", latex2rtf_bin_url, "-o", "binary.bin")
	err = latex2rtf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	latex2rtf_src_url := "https://downloads.sourceforge.net/project/latex2rtf/latex2rtf-unix/2.3.18/latex2rtf-2.3.18a.src.tar.gz"
	latex2rtf_cmd_src := exec.Command("curl", "-L", latex2rtf_src_url, "-o", "source.tar.gz")
	err = latex2rtf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	latex2rtf_cmd_direct := exec.Command("./binary")
	err = latex2rtf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

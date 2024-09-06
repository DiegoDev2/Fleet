package main

// Wtf - Translate common Internet acronyms
// Homepage: https://sourceforge.net/projects/bsdwtf/

import (
	"fmt"
	
	"os/exec"
)

func installWtf() {
	// Método 1: Descargar y extraer .tar.gz
	wtf_tar_url := "https://downloads.sourceforge.net/project/bsdwtf/wtf-20230906.tar.gz"
	wtf_cmd_tar := exec.Command("curl", "-L", wtf_tar_url, "-o", "package.tar.gz")
	err := wtf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wtf_zip_url := "https://downloads.sourceforge.net/project/bsdwtf/wtf-20230906.zip"
	wtf_cmd_zip := exec.Command("curl", "-L", wtf_zip_url, "-o", "package.zip")
	err = wtf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wtf_bin_url := "https://downloads.sourceforge.net/project/bsdwtf/wtf-20230906.bin"
	wtf_cmd_bin := exec.Command("curl", "-L", wtf_bin_url, "-o", "binary.bin")
	err = wtf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wtf_src_url := "https://downloads.sourceforge.net/project/bsdwtf/wtf-20230906.src.tar.gz"
	wtf_cmd_src := exec.Command("curl", "-L", wtf_src_url, "-o", "source.tar.gz")
	err = wtf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wtf_cmd_direct := exec.Command("./binary")
	err = wtf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

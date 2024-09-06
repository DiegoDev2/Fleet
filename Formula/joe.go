package main

// Joe - Full featured terminal-based screen editor
// Homepage: https://joe-editor.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installJoe() {
	// Método 1: Descargar y extraer .tar.gz
	joe_tar_url := "https://downloads.sourceforge.net/project/joe-editor/JOE%20sources/joe-4.6/joe-4.6.tar.gz"
	joe_cmd_tar := exec.Command("curl", "-L", joe_tar_url, "-o", "package.tar.gz")
	err := joe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	joe_zip_url := "https://downloads.sourceforge.net/project/joe-editor/JOE%20sources/joe-4.6/joe-4.6.zip"
	joe_cmd_zip := exec.Command("curl", "-L", joe_zip_url, "-o", "package.zip")
	err = joe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	joe_bin_url := "https://downloads.sourceforge.net/project/joe-editor/JOE%20sources/joe-4.6/joe-4.6.bin"
	joe_cmd_bin := exec.Command("curl", "-L", joe_bin_url, "-o", "binary.bin")
	err = joe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	joe_src_url := "https://downloads.sourceforge.net/project/joe-editor/JOE%20sources/joe-4.6/joe-4.6.src.tar.gz"
	joe_cmd_src := exec.Command("curl", "-L", joe_src_url, "-o", "source.tar.gz")
	err = joe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	joe_cmd_direct := exec.Command("./binary")
	err = joe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

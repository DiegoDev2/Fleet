package main

// Hilite - CLI tool that runs a command and highlights STDERR output
// Homepage: https://sourceforge.net/projects/hilite/

import (
	"fmt"
	
	"os/exec"
)

func installHilite() {
	// Método 1: Descargar y extraer .tar.gz
	hilite_tar_url := "https://downloads.sourceforge.net/project/hilite/hilite/1.5/hilite.c"
	hilite_cmd_tar := exec.Command("curl", "-L", hilite_tar_url, "-o", "package.tar.gz")
	err := hilite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hilite_zip_url := "https://downloads.sourceforge.net/project/hilite/hilite/1.5/hilite.c"
	hilite_cmd_zip := exec.Command("curl", "-L", hilite_zip_url, "-o", "package.zip")
	err = hilite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hilite_bin_url := "https://downloads.sourceforge.net/project/hilite/hilite/1.5/hilite.c"
	hilite_cmd_bin := exec.Command("curl", "-L", hilite_bin_url, "-o", "binary.bin")
	err = hilite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hilite_src_url := "https://downloads.sourceforge.net/project/hilite/hilite/1.5/hilite.c"
	hilite_cmd_src := exec.Command("curl", "-L", hilite_src_url, "-o", "source.tar.gz")
	err = hilite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hilite_cmd_direct := exec.Command("./binary")
	err = hilite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

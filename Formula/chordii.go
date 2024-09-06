package main

// Chordii - Text file to music sheet converter
// Homepage: https://www.vromans.org/johan/projects/Chordii/

import (
	"fmt"
	
	"os/exec"
)

func installChordii() {
	// Método 1: Descargar y extraer .tar.gz
	chordii_tar_url := "https://downloads.sourceforge.net/project/chordii/chordii/4.5/chordii-4.5.3b.tar.gz"
	chordii_cmd_tar := exec.Command("curl", "-L", chordii_tar_url, "-o", "package.tar.gz")
	err := chordii_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chordii_zip_url := "https://downloads.sourceforge.net/project/chordii/chordii/4.5/chordii-4.5.3b.zip"
	chordii_cmd_zip := exec.Command("curl", "-L", chordii_zip_url, "-o", "package.zip")
	err = chordii_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chordii_bin_url := "https://downloads.sourceforge.net/project/chordii/chordii/4.5/chordii-4.5.3b.bin"
	chordii_cmd_bin := exec.Command("curl", "-L", chordii_bin_url, "-o", "binary.bin")
	err = chordii_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chordii_src_url := "https://downloads.sourceforge.net/project/chordii/chordii/4.5/chordii-4.5.3b.src.tar.gz"
	chordii_cmd_src := exec.Command("curl", "-L", chordii_src_url, "-o", "source.tar.gz")
	err = chordii_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chordii_cmd_direct := exec.Command("./binary")
	err = chordii_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

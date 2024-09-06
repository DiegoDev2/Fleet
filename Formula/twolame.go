package main

// TwoLame - Optimized MPEG Audio Layer 2 (MP2) encoder
// Homepage: https://www.twolame.org/

import (
	"fmt"
	
	"os/exec"
)

func installTwoLame() {
	// Método 1: Descargar y extraer .tar.gz
	twolame_tar_url := "https://downloads.sourceforge.net/project/twolame/twolame/0.4.0/twolame-0.4.0.tar.gz"
	twolame_cmd_tar := exec.Command("curl", "-L", twolame_tar_url, "-o", "package.tar.gz")
	err := twolame_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	twolame_zip_url := "https://downloads.sourceforge.net/project/twolame/twolame/0.4.0/twolame-0.4.0.zip"
	twolame_cmd_zip := exec.Command("curl", "-L", twolame_zip_url, "-o", "package.zip")
	err = twolame_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	twolame_bin_url := "https://downloads.sourceforge.net/project/twolame/twolame/0.4.0/twolame-0.4.0.bin"
	twolame_cmd_bin := exec.Command("curl", "-L", twolame_bin_url, "-o", "binary.bin")
	err = twolame_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	twolame_src_url := "https://downloads.sourceforge.net/project/twolame/twolame/0.4.0/twolame-0.4.0.src.tar.gz"
	twolame_cmd_src := exec.Command("curl", "-L", twolame_src_url, "-o", "source.tar.gz")
	err = twolame_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	twolame_cmd_direct := exec.Command("./binary")
	err = twolame_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

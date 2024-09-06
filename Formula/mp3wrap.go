package main

// Mp3wrap - Wrap two or more mp3 files in a single large file
// Homepage: https://mp3wrap.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMp3wrap() {
	// Método 1: Descargar y extraer .tar.gz
	mp3wrap_tar_url := "https://downloads.sourceforge.net/project/mp3wrap/mp3wrap/mp3wrap%200.5/mp3wrap-0.5-src.tar.gz"
	mp3wrap_cmd_tar := exec.Command("curl", "-L", mp3wrap_tar_url, "-o", "package.tar.gz")
	err := mp3wrap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3wrap_zip_url := "https://downloads.sourceforge.net/project/mp3wrap/mp3wrap/mp3wrap%200.5/mp3wrap-0.5-src.zip"
	mp3wrap_cmd_zip := exec.Command("curl", "-L", mp3wrap_zip_url, "-o", "package.zip")
	err = mp3wrap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3wrap_bin_url := "https://downloads.sourceforge.net/project/mp3wrap/mp3wrap/mp3wrap%200.5/mp3wrap-0.5-src.bin"
	mp3wrap_cmd_bin := exec.Command("curl", "-L", mp3wrap_bin_url, "-o", "binary.bin")
	err = mp3wrap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3wrap_src_url := "https://downloads.sourceforge.net/project/mp3wrap/mp3wrap/mp3wrap%200.5/mp3wrap-0.5-src.src.tar.gz"
	mp3wrap_cmd_src := exec.Command("curl", "-L", mp3wrap_src_url, "-o", "source.tar.gz")
	err = mp3wrap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3wrap_cmd_direct := exec.Command("./binary")
	err = mp3wrap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

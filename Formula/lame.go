package main

// Lame - High quality MPEG Audio Layer III (MP3) encoder
// Homepage: https://lame.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installLame() {
	// Método 1: Descargar y extraer .tar.gz
	lame_tar_url := "https://downloads.sourceforge.net/project/lame/lame/3.100/lame-3.100.tar.gz"
	lame_cmd_tar := exec.Command("curl", "-L", lame_tar_url, "-o", "package.tar.gz")
	err := lame_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lame_zip_url := "https://downloads.sourceforge.net/project/lame/lame/3.100/lame-3.100.zip"
	lame_cmd_zip := exec.Command("curl", "-L", lame_zip_url, "-o", "package.zip")
	err = lame_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lame_bin_url := "https://downloads.sourceforge.net/project/lame/lame/3.100/lame-3.100.bin"
	lame_cmd_bin := exec.Command("curl", "-L", lame_bin_url, "-o", "binary.bin")
	err = lame_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lame_src_url := "https://downloads.sourceforge.net/project/lame/lame/3.100/lame-3.100.src.tar.gz"
	lame_cmd_src := exec.Command("curl", "-L", lame_src_url, "-o", "source.tar.gz")
	err = lame_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lame_cmd_direct := exec.Command("./binary")
	err = lame_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// Libcmph - C minimal perfect hashing library
// Homepage: https://cmph.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibcmph() {
	// Método 1: Descargar y extraer .tar.gz
	libcmph_tar_url := "https://downloads.sourceforge.net/project/cmph/v2.0.2/cmph-2.0.2.tar.gz"
	libcmph_cmd_tar := exec.Command("curl", "-L", libcmph_tar_url, "-o", "package.tar.gz")
	err := libcmph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcmph_zip_url := "https://downloads.sourceforge.net/project/cmph/v2.0.2/cmph-2.0.2.zip"
	libcmph_cmd_zip := exec.Command("curl", "-L", libcmph_zip_url, "-o", "package.zip")
	err = libcmph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcmph_bin_url := "https://downloads.sourceforge.net/project/cmph/v2.0.2/cmph-2.0.2.bin"
	libcmph_cmd_bin := exec.Command("curl", "-L", libcmph_bin_url, "-o", "binary.bin")
	err = libcmph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcmph_src_url := "https://downloads.sourceforge.net/project/cmph/v2.0.2/cmph-2.0.2.src.tar.gz"
	libcmph_cmd_src := exec.Command("curl", "-L", libcmph_src_url, "-o", "source.tar.gz")
	err = libcmph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcmph_cmd_direct := exec.Command("./binary")
	err = libcmph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

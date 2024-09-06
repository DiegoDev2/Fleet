package main

// LibxmpLite - Lite libxmp
// Homepage: https://xmp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibxmpLite() {
	// Método 1: Descargar y extraer .tar.gz
	libxmplite_tar_url := "https://downloads.sourceforge.net/project/xmp/libxmp/4.6.0/libxmp-lite-4.6.0.tar.gz"
	libxmplite_cmd_tar := exec.Command("curl", "-L", libxmplite_tar_url, "-o", "package.tar.gz")
	err := libxmplite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxmplite_zip_url := "https://downloads.sourceforge.net/project/xmp/libxmp/4.6.0/libxmp-lite-4.6.0.zip"
	libxmplite_cmd_zip := exec.Command("curl", "-L", libxmplite_zip_url, "-o", "package.zip")
	err = libxmplite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxmplite_bin_url := "https://downloads.sourceforge.net/project/xmp/libxmp/4.6.0/libxmp-lite-4.6.0.bin"
	libxmplite_cmd_bin := exec.Command("curl", "-L", libxmplite_bin_url, "-o", "binary.bin")
	err = libxmplite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxmplite_src_url := "https://downloads.sourceforge.net/project/xmp/libxmp/4.6.0/libxmp-lite-4.6.0.src.tar.gz"
	libxmplite_cmd_src := exec.Command("curl", "-L", libxmplite_src_url, "-o", "source.tar.gz")
	err = libxmplite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxmplite_cmd_direct := exec.Command("./binary")
	err = libxmplite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// LibvoAacenc - VisualOn AAC encoder library
// Homepage: https://opencore-amr.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibvoAacenc() {
	// Método 1: Descargar y extraer .tar.gz
	libvoaacenc_tar_url := "https://downloads.sourceforge.net/project/opencore-amr/vo-aacenc/vo-aacenc-0.1.3.tar.gz"
	libvoaacenc_cmd_tar := exec.Command("curl", "-L", libvoaacenc_tar_url, "-o", "package.tar.gz")
	err := libvoaacenc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvoaacenc_zip_url := "https://downloads.sourceforge.net/project/opencore-amr/vo-aacenc/vo-aacenc-0.1.3.zip"
	libvoaacenc_cmd_zip := exec.Command("curl", "-L", libvoaacenc_zip_url, "-o", "package.zip")
	err = libvoaacenc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvoaacenc_bin_url := "https://downloads.sourceforge.net/project/opencore-amr/vo-aacenc/vo-aacenc-0.1.3.bin"
	libvoaacenc_cmd_bin := exec.Command("curl", "-L", libvoaacenc_bin_url, "-o", "binary.bin")
	err = libvoaacenc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvoaacenc_src_url := "https://downloads.sourceforge.net/project/opencore-amr/vo-aacenc/vo-aacenc-0.1.3.src.tar.gz"
	libvoaacenc_cmd_src := exec.Command("curl", "-L", libvoaacenc_src_url, "-o", "source.tar.gz")
	err = libvoaacenc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvoaacenc_cmd_direct := exec.Command("./binary")
	err = libvoaacenc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// Libxmp - C library for playback of module music (MOD, S3M, IT, etc)
// Homepage: https://xmp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibxmp() {
	// Método 1: Descargar y extraer .tar.gz
	libxmp_tar_url := "https://downloads.sourceforge.net/project/xmp/libxmp/4.6.0/libxmp-4.6.0.tar.gz"
	libxmp_cmd_tar := exec.Command("curl", "-L", libxmp_tar_url, "-o", "package.tar.gz")
	err := libxmp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxmp_zip_url := "https://downloads.sourceforge.net/project/xmp/libxmp/4.6.0/libxmp-4.6.0.zip"
	libxmp_cmd_zip := exec.Command("curl", "-L", libxmp_zip_url, "-o", "package.zip")
	err = libxmp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxmp_bin_url := "https://downloads.sourceforge.net/project/xmp/libxmp/4.6.0/libxmp-4.6.0.bin"
	libxmp_cmd_bin := exec.Command("curl", "-L", libxmp_bin_url, "-o", "binary.bin")
	err = libxmp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxmp_src_url := "https://downloads.sourceforge.net/project/xmp/libxmp/4.6.0/libxmp-4.6.0.src.tar.gz"
	libxmp_cmd_src := exec.Command("curl", "-L", libxmp_src_url, "-o", "source.tar.gz")
	err = libxmp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxmp_cmd_direct := exec.Command("./binary")
	err = libxmp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
}

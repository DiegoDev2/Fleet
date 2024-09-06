package main

// Libogg - Ogg Bitstream Library
// Homepage: https://www.xiph.org/ogg/

import (
	"fmt"
	
	"os/exec"
)

func installLibogg() {
	// Método 1: Descargar y extraer .tar.gz
	libogg_tar_url := "https://ftp.osuosl.org/pub/xiph/releases/ogg/libogg-1.3.5.tar.gz"
	libogg_cmd_tar := exec.Command("curl", "-L", libogg_tar_url, "-o", "package.tar.gz")
	err := libogg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libogg_zip_url := "https://ftp.osuosl.org/pub/xiph/releases/ogg/libogg-1.3.5.zip"
	libogg_cmd_zip := exec.Command("curl", "-L", libogg_zip_url, "-o", "package.zip")
	err = libogg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libogg_bin_url := "https://ftp.osuosl.org/pub/xiph/releases/ogg/libogg-1.3.5.bin"
	libogg_cmd_bin := exec.Command("curl", "-L", libogg_bin_url, "-o", "binary.bin")
	err = libogg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libogg_src_url := "https://ftp.osuosl.org/pub/xiph/releases/ogg/libogg-1.3.5.src.tar.gz"
	libogg_cmd_src := exec.Command("curl", "-L", libogg_src_url, "-o", "source.tar.gz")
	err = libogg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libogg_cmd_direct := exec.Command("./binary")
	err = libogg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

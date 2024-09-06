package main

// Codec2 - Open source speech codec
// Homepage: https://www.rowetel.com/?page_id=452

import (
	"fmt"
	
	"os/exec"
)

func installCodec2() {
	// Método 1: Descargar y extraer .tar.gz
	codec2_tar_url := "https://github.com/drowe67/codec2/archive/refs/tags/1.2.0.tar.gz"
	codec2_cmd_tar := exec.Command("curl", "-L", codec2_tar_url, "-o", "package.tar.gz")
	err := codec2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codec2_zip_url := "https://github.com/drowe67/codec2/archive/refs/tags/1.2.0.zip"
	codec2_cmd_zip := exec.Command("curl", "-L", codec2_zip_url, "-o", "package.zip")
	err = codec2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codec2_bin_url := "https://github.com/drowe67/codec2/archive/refs/tags/1.2.0.bin"
	codec2_cmd_bin := exec.Command("curl", "-L", codec2_bin_url, "-o", "binary.bin")
	err = codec2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codec2_src_url := "https://github.com/drowe67/codec2/archive/refs/tags/1.2.0.src.tar.gz"
	codec2_cmd_src := exec.Command("curl", "-L", codec2_src_url, "-o", "source.tar.gz")
	err = codec2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codec2_cmd_direct := exec.Command("./binary")
	err = codec2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

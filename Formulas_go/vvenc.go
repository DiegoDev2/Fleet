package main

// Vvenc - Fraunhofer Versatile Video Encoder
// Homepage: https://github.com/fraunhoferhhi/vvenc

import (
	"fmt"
	
	"os/exec"
)

func installVvenc() {
	// Método 1: Descargar y extraer .tar.gz
	vvenc_tar_url := "https://github.com/fraunhoferhhi/vvenc/archive/refs/tags/v1.12.0.tar.gz"
	vvenc_cmd_tar := exec.Command("curl", "-L", vvenc_tar_url, "-o", "package.tar.gz")
	err := vvenc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vvenc_zip_url := "https://github.com/fraunhoferhhi/vvenc/archive/refs/tags/v1.12.0.zip"
	vvenc_cmd_zip := exec.Command("curl", "-L", vvenc_zip_url, "-o", "package.zip")
	err = vvenc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vvenc_bin_url := "https://github.com/fraunhoferhhi/vvenc/archive/refs/tags/v1.12.0.bin"
	vvenc_cmd_bin := exec.Command("curl", "-L", vvenc_bin_url, "-o", "binary.bin")
	err = vvenc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vvenc_src_url := "https://github.com/fraunhoferhhi/vvenc/archive/refs/tags/v1.12.0.src.tar.gz"
	vvenc_cmd_src := exec.Command("curl", "-L", vvenc_src_url, "-o", "source.tar.gz")
	err = vvenc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vvenc_cmd_direct := exec.Command("./binary")
	err = vvenc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

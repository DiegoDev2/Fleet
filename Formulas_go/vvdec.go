package main

// Vvdec - Fraunhofer Versatile Video Decoder
// Homepage: https://github.com/fraunhoferhhi/vvdec

import (
	"fmt"
	
	"os/exec"
)

func installVvdec() {
	// Método 1: Descargar y extraer .tar.gz
	vvdec_tar_url := "https://github.com/fraunhoferhhi/vvdec/archive/refs/tags/v2.3.0.tar.gz"
	vvdec_cmd_tar := exec.Command("curl", "-L", vvdec_tar_url, "-o", "package.tar.gz")
	err := vvdec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vvdec_zip_url := "https://github.com/fraunhoferhhi/vvdec/archive/refs/tags/v2.3.0.zip"
	vvdec_cmd_zip := exec.Command("curl", "-L", vvdec_zip_url, "-o", "package.zip")
	err = vvdec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vvdec_bin_url := "https://github.com/fraunhoferhhi/vvdec/archive/refs/tags/v2.3.0.bin"
	vvdec_cmd_bin := exec.Command("curl", "-L", vvdec_bin_url, "-o", "binary.bin")
	err = vvdec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vvdec_src_url := "https://github.com/fraunhoferhhi/vvdec/archive/refs/tags/v2.3.0.src.tar.gz"
	vvdec_cmd_src := exec.Command("curl", "-L", vvdec_src_url, "-o", "source.tar.gz")
	err = vvdec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vvdec_cmd_direct := exec.Command("./binary")
	err = vvdec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

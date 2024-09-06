package main

// SvtAv1 - AV1 encoder
// Homepage: https://gitlab.com/AOMediaCodec/SVT-AV1

import (
	"fmt"
	
	"os/exec"
)

func installSvtAv1() {
	// Método 1: Descargar y extraer .tar.gz
	svtav1_tar_url := "https://gitlab.com/AOMediaCodec/SVT-AV1/-/archive/v2.2.1/SVT-AV1-v2.2.1.tar.bz2"
	svtav1_cmd_tar := exec.Command("curl", "-L", svtav1_tar_url, "-o", "package.tar.gz")
	err := svtav1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	svtav1_zip_url := "https://gitlab.com/AOMediaCodec/SVT-AV1/-/archive/v2.2.1/SVT-AV1-v2.2.1.tar.bz2"
	svtav1_cmd_zip := exec.Command("curl", "-L", svtav1_zip_url, "-o", "package.zip")
	err = svtav1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	svtav1_bin_url := "https://gitlab.com/AOMediaCodec/SVT-AV1/-/archive/v2.2.1/SVT-AV1-v2.2.1.tar.bz2"
	svtav1_cmd_bin := exec.Command("curl", "-L", svtav1_bin_url, "-o", "binary.bin")
	err = svtav1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	svtav1_src_url := "https://gitlab.com/AOMediaCodec/SVT-AV1/-/archive/v2.2.1/SVT-AV1-v2.2.1.tar.bz2"
	svtav1_cmd_src := exec.Command("curl", "-L", svtav1_src_url, "-o", "source.tar.gz")
	err = svtav1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	svtav1_cmd_direct := exec.Command("./binary")
	err = svtav1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
}

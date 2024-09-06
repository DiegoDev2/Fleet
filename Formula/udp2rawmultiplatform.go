package main

// Udp2rawMultiplatform - Multi-platform(cross-platform) version of udp2raw-tunnel client
// Homepage: https://github.com/wangyu-/udp2raw-multiplatform

import (
	"fmt"
	
	"os/exec"
)

func installUdp2rawMultiplatform() {
	// Método 1: Descargar y extraer .tar.gz
	udp2rawmultiplatform_tar_url := "https://github.com/wangyu-/udp2raw-multiplatform/archive/refs/tags/20230206.0.tar.gz"
	udp2rawmultiplatform_cmd_tar := exec.Command("curl", "-L", udp2rawmultiplatform_tar_url, "-o", "package.tar.gz")
	err := udp2rawmultiplatform_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	udp2rawmultiplatform_zip_url := "https://github.com/wangyu-/udp2raw-multiplatform/archive/refs/tags/20230206.0.zip"
	udp2rawmultiplatform_cmd_zip := exec.Command("curl", "-L", udp2rawmultiplatform_zip_url, "-o", "package.zip")
	err = udp2rawmultiplatform_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	udp2rawmultiplatform_bin_url := "https://github.com/wangyu-/udp2raw-multiplatform/archive/refs/tags/20230206.0.bin"
	udp2rawmultiplatform_cmd_bin := exec.Command("curl", "-L", udp2rawmultiplatform_bin_url, "-o", "binary.bin")
	err = udp2rawmultiplatform_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	udp2rawmultiplatform_src_url := "https://github.com/wangyu-/udp2raw-multiplatform/archive/refs/tags/20230206.0.src.tar.gz"
	udp2rawmultiplatform_cmd_src := exec.Command("curl", "-L", udp2rawmultiplatform_src_url, "-o", "source.tar.gz")
	err = udp2rawmultiplatform_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	udp2rawmultiplatform_cmd_direct := exec.Command("./binary")
	err = udp2rawmultiplatform_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libnet")
	exec.Command("latte", "install", "libnet").Run()
}

package main

// CernNdiff - Numerical diff tool
// Homepage: https://mad.web.cern.ch/mad/

import (
	"fmt"
	
	"os/exec"
)

func installCernNdiff() {
	// Método 1: Descargar y extraer .tar.gz
	cernndiff_tar_url := "https://github.com/MethodicalAcceleratorDesign/MAD-X/archive/refs/tags/5.09.03.tar.gz"
	cernndiff_cmd_tar := exec.Command("curl", "-L", cernndiff_tar_url, "-o", "package.tar.gz")
	err := cernndiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cernndiff_zip_url := "https://github.com/MethodicalAcceleratorDesign/MAD-X/archive/refs/tags/5.09.03.zip"
	cernndiff_cmd_zip := exec.Command("curl", "-L", cernndiff_zip_url, "-o", "package.zip")
	err = cernndiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cernndiff_bin_url := "https://github.com/MethodicalAcceleratorDesign/MAD-X/archive/refs/tags/5.09.03.bin"
	cernndiff_cmd_bin := exec.Command("curl", "-L", cernndiff_bin_url, "-o", "binary.bin")
	err = cernndiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cernndiff_src_url := "https://github.com/MethodicalAcceleratorDesign/MAD-X/archive/refs/tags/5.09.03.src.tar.gz"
	cernndiff_cmd_src := exec.Command("curl", "-L", cernndiff_src_url, "-o", "source.tar.gz")
	err = cernndiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cernndiff_cmd_direct := exec.Command("./binary")
	err = cernndiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}

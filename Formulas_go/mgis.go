package main

// Mgis - Provide tools to handle MFront generic interface behaviours
// Homepage: https://thelfer.github.io/mgis/web/index.html

import (
	"fmt"
	
	"os/exec"
)

func installMgis() {
	// Método 1: Descargar y extraer .tar.gz
	mgis_tar_url := "https://github.com/thelfer/MFrontGenericInterfaceSupport/archive/refs/tags/MFrontGenericInterfaceSupport-2.2.tar.gz"
	mgis_cmd_tar := exec.Command("curl", "-L", mgis_tar_url, "-o", "package.tar.gz")
	err := mgis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mgis_zip_url := "https://github.com/thelfer/MFrontGenericInterfaceSupport/archive/refs/tags/MFrontGenericInterfaceSupport-2.2.zip"
	mgis_cmd_zip := exec.Command("curl", "-L", mgis_zip_url, "-o", "package.zip")
	err = mgis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mgis_bin_url := "https://github.com/thelfer/MFrontGenericInterfaceSupport/archive/refs/tags/MFrontGenericInterfaceSupport-2.2.bin"
	mgis_cmd_bin := exec.Command("curl", "-L", mgis_bin_url, "-o", "binary.bin")
	err = mgis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mgis_src_url := "https://github.com/thelfer/MFrontGenericInterfaceSupport/archive/refs/tags/MFrontGenericInterfaceSupport-2.2.src.tar.gz"
	mgis_cmd_src := exec.Command("curl", "-L", mgis_src_url, "-o", "source.tar.gz")
	err = mgis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mgis_cmd_direct := exec.Command("./binary")
	err = mgis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: boost-python3")
exec.Command("latte", "install", "boost-python3")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}

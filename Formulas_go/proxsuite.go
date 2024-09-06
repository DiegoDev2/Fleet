package main

// Proxsuite - Advanced Proximal Optimization Toolbox
// Homepage: https://github.com/Simple-Robotics/proxsuite

import (
	"fmt"
	
	"os/exec"
)

func installProxsuite() {
	// Método 1: Descargar y extraer .tar.gz
	proxsuite_tar_url := "https://github.com/Simple-Robotics/proxsuite/releases/download/v0.6.7/proxsuite-0.6.7.tar.gz"
	proxsuite_cmd_tar := exec.Command("curl", "-L", proxsuite_tar_url, "-o", "package.tar.gz")
	err := proxsuite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proxsuite_zip_url := "https://github.com/Simple-Robotics/proxsuite/releases/download/v0.6.7/proxsuite-0.6.7.zip"
	proxsuite_cmd_zip := exec.Command("curl", "-L", proxsuite_zip_url, "-o", "package.zip")
	err = proxsuite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proxsuite_bin_url := "https://github.com/Simple-Robotics/proxsuite/releases/download/v0.6.7/proxsuite-0.6.7.bin"
	proxsuite_cmd_bin := exec.Command("curl", "-L", proxsuite_bin_url, "-o", "binary.bin")
	err = proxsuite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proxsuite_src_url := "https://github.com/Simple-Robotics/proxsuite/releases/download/v0.6.7/proxsuite-0.6.7.src.tar.gz"
	proxsuite_cmd_src := exec.Command("curl", "-L", proxsuite_src_url, "-o", "source.tar.gz")
	err = proxsuite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proxsuite_cmd_direct := exec.Command("./binary")
	err = proxsuite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: scipy")
exec.Command("latte", "install", "scipy")
	fmt.Println("Instalando dependencia: simde")
exec.Command("latte", "install", "simde")
}

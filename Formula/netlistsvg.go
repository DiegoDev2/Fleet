package main

// Netlistsvg - Draws an SVG schematic from a yosys JSON netlist
// Homepage: https://github.com/nturley/netlistsvg

import (
	"fmt"
	
	"os/exec"
)

func installNetlistsvg() {
	// Método 1: Descargar y extraer .tar.gz
	netlistsvg_tar_url := "https://github.com/nturley/netlistsvg/archive/refs/tags/v1.0.2.tar.gz"
	netlistsvg_cmd_tar := exec.Command("curl", "-L", netlistsvg_tar_url, "-o", "package.tar.gz")
	err := netlistsvg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netlistsvg_zip_url := "https://github.com/nturley/netlistsvg/archive/refs/tags/v1.0.2.zip"
	netlistsvg_cmd_zip := exec.Command("curl", "-L", netlistsvg_zip_url, "-o", "package.zip")
	err = netlistsvg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netlistsvg_bin_url := "https://github.com/nturley/netlistsvg/archive/refs/tags/v1.0.2.bin"
	netlistsvg_cmd_bin := exec.Command("curl", "-L", netlistsvg_bin_url, "-o", "binary.bin")
	err = netlistsvg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netlistsvg_src_url := "https://github.com/nturley/netlistsvg/archive/refs/tags/v1.0.2.src.tar.gz"
	netlistsvg_cmd_src := exec.Command("curl", "-L", netlistsvg_src_url, "-o", "source.tar.gz")
	err = netlistsvg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netlistsvg_cmd_direct := exec.Command("./binary")
	err = netlistsvg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: yosys")
	exec.Command("latte", "install", "yosys").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}

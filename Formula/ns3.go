package main

// Ns3 - Discrete-event network simulator
// Homepage: https://www.nsnam.org/

import (
	"fmt"
	
	"os/exec"
)

func installNs3() {
	// Método 1: Descargar y extraer .tar.gz
	ns3_tar_url := "https://gitlab.com/nsnam/ns-3-dev/-/archive/ns-3.42/ns-3-dev-ns-3.42.tar.gz"
	ns3_cmd_tar := exec.Command("curl", "-L", ns3_tar_url, "-o", "package.tar.gz")
	err := ns3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ns3_zip_url := "https://gitlab.com/nsnam/ns-3-dev/-/archive/ns-3.42/ns-3-dev-ns-3.42.zip"
	ns3_cmd_zip := exec.Command("curl", "-L", ns3_zip_url, "-o", "package.zip")
	err = ns3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ns3_bin_url := "https://gitlab.com/nsnam/ns-3-dev/-/archive/ns-3.42/ns-3-dev-ns-3.42.bin"
	ns3_cmd_bin := exec.Command("curl", "-L", ns3_bin_url, "-o", "binary.bin")
	err = ns3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ns3_src_url := "https://gitlab.com/nsnam/ns-3-dev/-/archive/ns-3.42/ns-3-dev-ns-3.42.src.tar.gz"
	ns3_cmd_src := exec.Command("curl", "-L", ns3_src_url, "-o", "source.tar.gz")
	err = ns3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ns3_cmd_direct := exec.Command("./binary")
	err = ns3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gsl")
	exec.Command("latte", "install", "gsl").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
}

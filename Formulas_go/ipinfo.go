package main

// Ipinfo - Tool for calculation of IP networks
// Homepage: https://kyberdigi.cz/projects/ipinfo/

import (
	"fmt"
	
	"os/exec"
)

func installIpinfo() {
	// Método 1: Descargar y extraer .tar.gz
	ipinfo_tar_url := "https://kyberdigi.cz/projects/ipinfo/files/ipinfo-1.2.tar.gz"
	ipinfo_cmd_tar := exec.Command("curl", "-L", ipinfo_tar_url, "-o", "package.tar.gz")
	err := ipinfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipinfo_zip_url := "https://kyberdigi.cz/projects/ipinfo/files/ipinfo-1.2.zip"
	ipinfo_cmd_zip := exec.Command("curl", "-L", ipinfo_zip_url, "-o", "package.zip")
	err = ipinfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipinfo_bin_url := "https://kyberdigi.cz/projects/ipinfo/files/ipinfo-1.2.bin"
	ipinfo_cmd_bin := exec.Command("curl", "-L", ipinfo_bin_url, "-o", "binary.bin")
	err = ipinfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipinfo_src_url := "https://kyberdigi.cz/projects/ipinfo/files/ipinfo-1.2.src.tar.gz"
	ipinfo_cmd_src := exec.Command("curl", "-L", ipinfo_src_url, "-o", "source.tar.gz")
	err = ipinfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipinfo_cmd_direct := exec.Command("./binary")
	err = ipinfo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

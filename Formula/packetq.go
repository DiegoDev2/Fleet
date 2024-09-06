package main

// Packetq - SQL-like frontend to PCAP files
// Homepage: https://www.dns-oarc.net/tools/packetq

import (
	"fmt"
	
	"os/exec"
)

func installPacketq() {
	// Método 1: Descargar y extraer .tar.gz
	packetq_tar_url := "https://www.dns-oarc.net/files/packetq/packetq-1.7.3.tar.gz"
	packetq_cmd_tar := exec.Command("curl", "-L", packetq_tar_url, "-o", "package.tar.gz")
	err := packetq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	packetq_zip_url := "https://www.dns-oarc.net/files/packetq/packetq-1.7.3.zip"
	packetq_cmd_zip := exec.Command("curl", "-L", packetq_zip_url, "-o", "package.zip")
	err = packetq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	packetq_bin_url := "https://www.dns-oarc.net/files/packetq/packetq-1.7.3.bin"
	packetq_cmd_bin := exec.Command("curl", "-L", packetq_bin_url, "-o", "binary.bin")
	err = packetq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	packetq_src_url := "https://www.dns-oarc.net/files/packetq/packetq-1.7.3.src.tar.gz"
	packetq_cmd_src := exec.Command("curl", "-L", packetq_src_url, "-o", "source.tar.gz")
	err = packetq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	packetq_cmd_direct := exec.Command("./binary")
	err = packetq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

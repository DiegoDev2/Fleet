package main

// SpiceProtocol - Headers for SPICE protocol
// Homepage: https://www.spice-space.org/

import (
	"fmt"
	
	"os/exec"
)

func installSpiceProtocol() {
	// Método 1: Descargar y extraer .tar.gz
	spiceprotocol_tar_url := "https://www.spice-space.org/download/releases/spice-protocol-0.14.4.tar.xz"
	spiceprotocol_cmd_tar := exec.Command("curl", "-L", spiceprotocol_tar_url, "-o", "package.tar.gz")
	err := spiceprotocol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spiceprotocol_zip_url := "https://www.spice-space.org/download/releases/spice-protocol-0.14.4.tar.xz"
	spiceprotocol_cmd_zip := exec.Command("curl", "-L", spiceprotocol_zip_url, "-o", "package.zip")
	err = spiceprotocol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spiceprotocol_bin_url := "https://www.spice-space.org/download/releases/spice-protocol-0.14.4.tar.xz"
	spiceprotocol_cmd_bin := exec.Command("curl", "-L", spiceprotocol_bin_url, "-o", "binary.bin")
	err = spiceprotocol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spiceprotocol_src_url := "https://www.spice-space.org/download/releases/spice-protocol-0.14.4.tar.xz"
	spiceprotocol_cmd_src := exec.Command("curl", "-L", spiceprotocol_src_url, "-o", "source.tar.gz")
	err = spiceprotocol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spiceprotocol_cmd_direct := exec.Command("./binary")
	err = spiceprotocol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}

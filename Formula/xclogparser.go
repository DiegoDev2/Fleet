package main

// Xclogparser - Tool to parse the SLF serialization format used by Xcode
// Homepage: https://github.com/MobileNativeFoundation/XCLogParser

import (
	"fmt"
	
	"os/exec"
)

func installXclogparser() {
	// Método 1: Descargar y extraer .tar.gz
	xclogparser_tar_url := "https://github.com/MobileNativeFoundation/XCLogParser/archive/refs/tags/v0.2.39.tar.gz"
	xclogparser_cmd_tar := exec.Command("curl", "-L", xclogparser_tar_url, "-o", "package.tar.gz")
	err := xclogparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xclogparser_zip_url := "https://github.com/MobileNativeFoundation/XCLogParser/archive/refs/tags/v0.2.39.zip"
	xclogparser_cmd_zip := exec.Command("curl", "-L", xclogparser_zip_url, "-o", "package.zip")
	err = xclogparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xclogparser_bin_url := "https://github.com/MobileNativeFoundation/XCLogParser/archive/refs/tags/v0.2.39.bin"
	xclogparser_cmd_bin := exec.Command("curl", "-L", xclogparser_bin_url, "-o", "binary.bin")
	err = xclogparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xclogparser_src_url := "https://github.com/MobileNativeFoundation/XCLogParser/archive/refs/tags/v0.2.39.src.tar.gz"
	xclogparser_cmd_src := exec.Command("curl", "-L", xclogparser_src_url, "-o", "source.tar.gz")
	err = xclogparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xclogparser_cmd_direct := exec.Command("./binary")
	err = xclogparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

package main

// StripNondeterminism - Tool for stripping bits of non-deterministic information from files
// Homepage: https://salsa.debian.org/reproducible-builds/strip-nondeterminism

import (
	"fmt"
	
	"os/exec"
)

func installStripNondeterminism() {
	// Método 1: Descargar y extraer .tar.gz
	stripnondeterminism_tar_url := "https://salsa.debian.org/reproducible-builds/strip-nondeterminism/-/archive/1.14.0/strip-nondeterminism-1.14.0.tar.bz2"
	stripnondeterminism_cmd_tar := exec.Command("curl", "-L", stripnondeterminism_tar_url, "-o", "package.tar.gz")
	err := stripnondeterminism_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stripnondeterminism_zip_url := "https://salsa.debian.org/reproducible-builds/strip-nondeterminism/-/archive/1.14.0/strip-nondeterminism-1.14.0.tar.bz2"
	stripnondeterminism_cmd_zip := exec.Command("curl", "-L", stripnondeterminism_zip_url, "-o", "package.zip")
	err = stripnondeterminism_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stripnondeterminism_bin_url := "https://salsa.debian.org/reproducible-builds/strip-nondeterminism/-/archive/1.14.0/strip-nondeterminism-1.14.0.tar.bz2"
	stripnondeterminism_cmd_bin := exec.Command("curl", "-L", stripnondeterminism_bin_url, "-o", "binary.bin")
	err = stripnondeterminism_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stripnondeterminism_src_url := "https://salsa.debian.org/reproducible-builds/strip-nondeterminism/-/archive/1.14.0/strip-nondeterminism-1.14.0.tar.bz2"
	stripnondeterminism_cmd_src := exec.Command("curl", "-L", stripnondeterminism_src_url, "-o", "source.tar.gz")
	err = stripnondeterminism_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stripnondeterminism_cmd_direct := exec.Command("./binary")
	err = stripnondeterminism_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

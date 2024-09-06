package main

// Auditwheel - Auditing and relabeling cross-distribution Linux wheels
// Homepage: https://github.com/pypa/auditwheel

import (
	"fmt"
	
	"os/exec"
)

func installAuditwheel() {
	// Método 1: Descargar y extraer .tar.gz
	auditwheel_tar_url := "https://files.pythonhosted.org/packages/ce/12/af339761d80296e76033669e4179883108d8a7d79bb032bd58427a2b4485/auditwheel-6.1.0.tar.gz"
	auditwheel_cmd_tar := exec.Command("curl", "-L", auditwheel_tar_url, "-o", "package.tar.gz")
	err := auditwheel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	auditwheel_zip_url := "https://files.pythonhosted.org/packages/ce/12/af339761d80296e76033669e4179883108d8a7d79bb032bd58427a2b4485/auditwheel-6.1.0.zip"
	auditwheel_cmd_zip := exec.Command("curl", "-L", auditwheel_zip_url, "-o", "package.zip")
	err = auditwheel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	auditwheel_bin_url := "https://files.pythonhosted.org/packages/ce/12/af339761d80296e76033669e4179883108d8a7d79bb032bd58427a2b4485/auditwheel-6.1.0.bin"
	auditwheel_cmd_bin := exec.Command("curl", "-L", auditwheel_bin_url, "-o", "binary.bin")
	err = auditwheel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	auditwheel_src_url := "https://files.pythonhosted.org/packages/ce/12/af339761d80296e76033669e4179883108d8a7d79bb032bd58427a2b4485/auditwheel-6.1.0.src.tar.gz"
	auditwheel_cmd_src := exec.Command("curl", "-L", auditwheel_src_url, "-o", "source.tar.gz")
	err = auditwheel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	auditwheel_cmd_direct := exec.Command("./binary")
	err = auditwheel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}

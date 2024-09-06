package main

// PyqtBuilder - Tool to build PyQt
// Homepage: https://pyqt-builder.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installPyqtBuilder() {
	// Método 1: Descargar y extraer .tar.gz
	pyqtbuilder_tar_url := "https://files.pythonhosted.org/packages/e6/f5/daead7fd8ef3675ce55f4ef66dbe3287b0bdd74315f6b5a57718a020570b/pyqt_builder-1.16.4.tar.gz"
	pyqtbuilder_cmd_tar := exec.Command("curl", "-L", pyqtbuilder_tar_url, "-o", "package.tar.gz")
	err := pyqtbuilder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyqtbuilder_zip_url := "https://files.pythonhosted.org/packages/e6/f5/daead7fd8ef3675ce55f4ef66dbe3287b0bdd74315f6b5a57718a020570b/pyqt_builder-1.16.4.zip"
	pyqtbuilder_cmd_zip := exec.Command("curl", "-L", pyqtbuilder_zip_url, "-o", "package.zip")
	err = pyqtbuilder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyqtbuilder_bin_url := "https://files.pythonhosted.org/packages/e6/f5/daead7fd8ef3675ce55f4ef66dbe3287b0bdd74315f6b5a57718a020570b/pyqt_builder-1.16.4.bin"
	pyqtbuilder_cmd_bin := exec.Command("curl", "-L", pyqtbuilder_bin_url, "-o", "binary.bin")
	err = pyqtbuilder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyqtbuilder_src_url := "https://files.pythonhosted.org/packages/e6/f5/daead7fd8ef3675ce55f4ef66dbe3287b0bdd74315f6b5a57718a020570b/pyqt_builder-1.16.4.src.tar.gz"
	pyqtbuilder_cmd_src := exec.Command("curl", "-L", pyqtbuilder_src_url, "-o", "source.tar.gz")
	err = pyqtbuilder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyqtbuilder_cmd_direct := exec.Command("./binary")
	err = pyqtbuilder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
